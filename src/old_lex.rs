use std::collections::HashMap;
use std::fmt;

use fmt::Formatter;

lazy_static! {
    static ref KEYWORDS: HashMap<&'static str, TokenType> = {
        let mut keywords = HashMap::new();
        keywords.insert("f", TokenType::FUNCTION);
        keywords.insert("value", TokenType::VALUE);
        keywords.insert("true", TokenType::TRUE);
        keywords.insert("false", TokenType::FALSE);
        keywords.insert("if", TokenType::IF);
        keywords.insert("else", TokenType::ELSE);
        keywords.insert("return", TokenType::RETURN);
        keywords.insert("is", TokenType::EQ);
        keywords.insert("not", TokenType::NOT_EQ);
        keywords.insert("end", TokenType::CLOSEBLOCK);
        keywords.insert("update", TokenType::UPDATE);
        keywords.insert("constant", TokenType::CONSTANT);
        keywords.insert("and", TokenType::AND);
        keywords.insert("or", TokenType::OR);
        keywords
    };
}

#[derive(Debug)]
pub enum TokenType {
    // Endings
    ILLEGAL,
    EOF,
    // Identifiers + literals
    IDENT,
    INT(usize),
    // Operators
    ASSIGN,
    PLUS,
    EQ,
    NOT_EQ,

    // Delimiters
    COMMA,
    SEMICOLON,
    LPAREN,
    RPAREN,
    LBRACE,
    RBRACE,
    LBRACKET,
    RBRACKET,
    COLON,
    // Keywords
    FUNCTION,
    VALUE,
    UPDATE,
    CONSTANT,
    TRUE,
    FALSE,
    IF,
    ELSE,
    RETURN,
    STRING(String),
    MINUS,
    BANG,
    ASTERISK,
    SLASH,
    LT,
    GT,
    LTEQ,
    GTEQ,
    INCREMENT,
    DECREMENT,
    MINUSEQ,
    PLUSEQ,
    MULEQ,
    SLASHEQ,
    OR,
    AND,
    OPENBLOCK,
    CLOSEBLOCK,
}

use TokenType::*;

pub struct Token<'a> {
    Type: TokenType,
    literal: &'a str,
    line: usize,
    column: usize,

}

impl<'a> fmt::Display for Token<'a> {
    fn fmt(&self, f: &mut Formatter<'_>) -> fmt::Result {
        write!(f, "({}, {}) [{:?}] lit: '{}'", self.line, self.column, self.Type, self.literal)
    }
}

pub fn lookup_ident(identifier: &str) -> Option<&'static TokenType> {
   KEYWORDS.get(identifier)
}

pub struct LexerError {
    line: usize,
    column: usize,
    reason: String,
}
pub struct Lexer {
    pub input: String,
    pub position: usize,
    pub reading_position: usize,
    pub column: usize,
    pub line: usize,
    pub current_char: char,
    pub errors: Vec<LexerError>
}

impl Lexer {
    fn next_token(&mut self) -> Result<Token, String> {
        let t: Token;

        self.skipWhitespace();

        match self.current_char {
            '(' => t = self.newCharToken(LPAREN),
            ')' => t = self.newCharToken(RPAREN),
            '+' => t = self.newCharToken(PLUS),
            '-' => {
                if self.peek_character() == '>' {
                    t = self.newCharToken(OPENBLOCK)
                } else {

                }
            }
            ',' => t = self.newCharToken(COMMA),
            '{' => t = self.newCharToken(LBRACKET),
            _ => {
                if self.current_char.is_numeric() {
                    let int_str = self.consumeInteger();
                    let int = int_str.parse::<usize>().unwrap();

                    t.literal = int_str;
                    t.Type = INT(int);
                } else if self.current_char.is_alphabetic() {
                    let start_col = self.column;
                    let start_line = self.line;
                    let str_literal = self.consumeIdentifier().clone();
                    let TT = lookup_ident(str_literal);
                    match TT {
                        Some(&TT) => {
                            t = Token {
                                line: start_line,
                                column: start_col,
                                literal: str_literal,
                                Type: TT
                            }
                        },
                        None => {
                            t = Token {
                                line: start_line,
                                column: start_col,
                                literal: str_literal,
                                Type: IDENT
                            }
                        }
                    };
                } else if self.current_char.is_numeric() {
                    let start_col = self.column;
                    let start_line = self.line;
                    let str_literal = self.consumeInteger().clone();
                    let int = str_literal.parse::<usize>();

                    match int {
                        Ok(i) => t = Token {
                            line: start_line,
                            column: start_col,
                            literal: str_literal,
                            Type: INT(i)
                        },
                        Err(z) => {
                            self.add_error(format!("({}, {}) - {}", self.line, self.column, z));
                            return Result::Err(format!("({}, {}) - {}", self.line, self.column, z));
                        }
                    }
                    
                        
                } else {
                    self.add_error(format!("({}, {}) - {}", self.line, self.column, "illegal token detected."));
                    return Result::Err(format!("({}, {}) - {}", self.line, self.column, "illegal token detected."));
                }
            }

        }

        self.consume_char();

        return Result::Ok(t)
    }

    pub fn consume_char(&mut self) {
        if self.reading_position >= self.input.len() {
            self.current_char = 0 as char;
        } else {
            self.current_char = self.input.chars().nth(self.reading_position).unwrap();
        }

        self.position += 1;
        self.reading_position += 1;
        self.column += 1;
    }

    fn skipWhitespace(&mut self) {
        while self.current_char.is_ascii_whitespace() {
            if self.current_char == '\n' {
                self.line += 1;
                self.column = 1;
            }
            self.consume_char();
        }
    }

    fn consumeInteger(&mut self) -> &str {
        let pos = self.position;
        while self.current_char.is_ascii_digit() {
            self.consume_char()
        }

        return &self.input[pos..self.position];
    }

    fn consumeIdentifier<'a>(&'a mut self) -> &str {
        let pos = self.position;
        while self.current_char.is_alphabetic() {
            self.consume_char()
        }

        return &self.input[pos..self.position];
    }

    fn readString(&mut self) -> &str {
        let pos = self.position;
        loop {
            self.consume_char();
            if self.current_char == '"' || self.current_char == 0 as char {
                break;
            }
        }

        return &self.input[pos..self.position];
    }

    fn peek_character(&self) -> char {
        if self.reading_position >= self.input.len() {
            return 0 as char;
        } else {
            return self.input.chars().nth(self.reading_position).unwrap();
        }
    }

    fn newCharToken<'a>(&'a self, Type: TokenType) -> Token<'a> {
        Token {
            column: self.column,
            line: self.line,
            literal: self.current_char.to_string().as_str(),
            Type: Type
        }
    }

    fn add_error(&mut self, reason: String) {
        let mut owned = "LEXER ERROR: ".to_owned();
        owned.push_str(&reason);
        let err = LexerError {
            reason: owned,
            line: self.line,
            column: self.column
        };

        self.errors.push(err);
    }

    fn lex(&mut self) {
        let tokens = Vec::<Token>::new();
        while self.position != self.input.len()-1 {
            let token = self.next_token();


        }
    }
}

fn New(input: String) -> Lexer {
    Lexer {
        input: input,
        position: 0,
        reading_position: 1,
        column: 1,
        line: 1,
        current_char: input.chars().nth(0).unwrap(),
        errors: vec![]
    }
}