#[derive(Debug, Clone, Copy)]
pub struct Position {
    column: usize,
    line: usize
}

#[derive(Debug, Clone)]
pub enum TokenType {
    //Associated 
    String(String),
    Integer(usize),
    Float(f64),
    Identifier(String),

    LBracket,
    RBracket,
    LParen,
    RParen,
    LBrace,
    RBrace,

    Ampersand,
    DoubleAmpersand,
    Equal,
    NotEqual,
    Bar,
    DoubleBar,
    Assign,

    True,
    False,
    Value

}

#[derive(Debug)]
pub struct Token {
    position: Position,
    kind: TokenType
}

#[derive(Debug, Clone)]
pub enum LexerError {
    UndefinedError(Position),
    ExpectPeek(Position, TokenType, TokenType)
}

#[derive(Debug, Clone)]
pub(crate) struct Lexer<'a> {
    input: &'a str,
    literal_position: Position,
    position: usize,
    errors: Vec<LexerError>
}

impl<'a> Lexer<'a> {

    fn consume_char(&mut self) {
        self.literal_position.column += 1;
        self.position += 1;


    }
    
    fn get_nth_char(&self, position: Option<usize>) ->  Option<char> {
        if let Some(pos) = position { //Check if pos exists
            return self.input.chars().nth(pos)
        } else {
            return self.input.chars().nth(self.position)
        }
    }

    fn peek_char(&self) -> Option<char> {
        self.get_nth_char(Some(self.position + 1))
    }

    pub fn single_char_tok(&mut self, kind: TokenType) -> Option<Token> {
        let current_position = self.literal_position;
        self.consume_char();
        Some(Token{
            kind: kind,
            position: current_position
        })
    }

    pub fn double_char_tok(&mut self, kind: TokenType) -> Option<Token> {
        let current_position = self.literal_position;
        self.consume_char();
        self.consume_char();
        Some(Token {
            kind: kind,
            position: current_position
        })
    }

    pub fn string(&mut self) -> Option<Token> {
        let current_position = self.literal_position;
        self.consume_char();
        let mut owned = String::new();
        while let Some(ch) = self.get_nth_char(None) {
            match ch {
                '"' => {
                    self.consume_char();
                    break;
                },
                '\n' => {
                    self.literal_position.column = 0;
                    self.literal_position.line += 1;
                    owned.push(ch);
                    self.consume_char();
                }
                _ => {
                    owned.push(ch);
                    self.consume_char();
                }
            }
        }

        Some(Token {
            kind: TokenType::String(owned),
            position: current_position  
        })
    }

    pub fn ident(&mut self) -> Option<Token> {
        let current_position = self.literal_position;
        let mut owned = String::new();

        while let Some(ch) = self.get_nth_char(None) {
            match ch {
                'A'..='Z' | 'a'..='z' | '_' => {
                    owned.push(ch);
                    self.consume_char();
                },
                _ => break
            }
        }

        let kind = match owned.as_str() {
            "true" => TokenType::True,
            "false" => TokenType::False,
            "value" => TokenType::Value,
            _ => TokenType::Identifier(owned)
        };

        Some(Token{
            kind: kind,
            position: current_position
        })
    }

    pub fn integer(&mut self) -> Option<Token> {
        enum NumberTypes {
            Integer,
            Float
        }
        let current_position = self.literal_position;
        let mut num = String::new();
        let mut num_type = NumberTypes::Integer;

        while let Some(ch) = self.get_nth_char(None) {
            match ch {
                '0'..='9' => {
                    num.push(ch);
                    self.consume_char();
                },
                '.' if matches!(self.peek_char(), Some('0'..='9')) => {
                    num_type = NumberTypes::Float;
                    num.push(ch);
                    self.consume_char();
                },
                _ => break
            }
        }

        match num_type {
            NumberTypes::Integer => {
                match num.parse::<usize>() {
                    Ok(n) => Some(Token {
                        kind: TokenType::Integer(n),
                        position: current_position,
                    }),
                    _ => {
                        self.add_error(LexerError::UndefinedError(current_position));
                        None
                    }
                }
            },
            NumberTypes::Float => {
                match num.parse::<f64>() {
                    Ok(n) => Some(Token {
                        kind: TokenType::Float(n),
                        position: current_position,
                    }),
                    _ => {
                        self.add_error(LexerError::UndefinedError(current_position));
                        None
                    }
                }
            }
        }


    }

    pub fn add_error(&mut self, kind: LexerError) -> Option<Token> {
        self.errors.push(kind);

        None
    }

    pub fn errors(self) -> Vec<LexerError> {
        return self.errors.clone();
    }

    pub fn new(input: &'a str) -> Lexer<'a> {
        Self {
            literal_position: Position {
                column: 0,
                line: 1, // Remember: lines aren't zero-indexed.
            },
            position: 0,
            input: input,
            errors: Vec::new()
        }
    }
}

impl<'a> Iterator for Lexer<'a> {
    type Item = Token;

    fn next(&mut self) -> Option<Self::Item> {
        let current_character = self.get_nth_char(None);
        let peek_character = self.peek_char();

        if current_character == None {
            return None
        }

        match current_character.unwrap() {
            '(' => self.single_char_tok(TokenType::LParen),
            ')' => self.single_char_tok(TokenType::RParen),
            '[' => self.single_char_tok(TokenType::LBracket),
            ']' => self.single_char_tok(TokenType::RBracket),
            '{' => self.single_char_tok(TokenType::LBrace),
            '}' => self.single_char_tok(TokenType::RBrace),

            '=' => {
                match peek_character {
                    Some('=') => self.double_char_tok(TokenType::Equal),
                    None | Some(' ') => self.single_char_tok(TokenType::Assign),
                    _ => self.add_error(LexerError::UndefinedError(self.literal_position))
                }
            },
            '\n' => {
                self.literal_position.column = 0;
                self.literal_position.line += 1;
                self.consume_char();
                self.next()
            },
            ' ' | '\t' | '\r' => {
                self.consume_char();
                self.next()
            },
            '0'..='9' => self.integer(),
            'A'..='Z' | 'a'..='z' | '_' => self.ident(),
            '"' => self.string(),
            _ => self.add_error(LexerError::UndefinedError(self.literal_position))
        }
    }
}

