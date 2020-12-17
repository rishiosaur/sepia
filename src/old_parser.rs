use crate::{
    ast::{Expression, Program, Statement},
    lexer::Token,
    util::variant_eq,
};

use crate::lexer::TokenType;

#[derive(Debug, Copy, Clone)]
pub enum Precedence {
    LOWEST = 0,
    AND = 1,
    OR = 2,
    EQUALS = 3,
    LESSGREATER = 4,
    SUM = 5,
    PRODUCT = 6,
    PREFIX = 7,
    CALL = 8,
    INDEX = 9,
}

pub fn get_precedence(kind: TokenType) -> Option<Precedence> {
    match kind {
        TokenType::Equal => Some(Precedence::EQUALS),
        TokenType::NotEqual => Some(Precedence::EQUALS),
        TokenType::LT => Some(Precedence::LESSGREATER),
        TokenType::GT => Some(Precedence::LESSGREATER),
        TokenType::LTEq => Some(Precedence::LESSGREATER),
        TokenType::GTEq => Some(Precedence::LESSGREATER),
        TokenType::Or => Some(Precedence::OR),
        TokenType::And => Some(Precedence::AND),
        TokenType::Plus => Some(Precedence::SUM),
        TokenType::Minus => Some(Precedence::SUM),
        TokenType::Asterisk => Some(Precedence::PRODUCT),
        TokenType::Slash => Some(Precedence::PRODUCT),
        TokenType::LParen => Some(Precedence::CALL),
        TokenType::LBracket => Some(Precedence::INDEX),
        _ => None,
    }
}

pub fn int_precedence(precedence: Precedence) -> usize {
    precedence as usize
}

pub struct Parser {
    pub tokens: Vec<Token>,
    pub position: usize,
}

impl Parser {
    pub fn current_token(&self) -> Option<Token> {
        self.tokens.get(self.position).cloned()
    }

    pub fn peek_token(&self) -> Option<Token> {
        self.tokens.get(self.position + 1).cloned()
    }

    pub fn current_precedence(&self) -> Precedence {
        match get_precedence(self.current_token().unwrap().clone().kind) {
            Some(p) => p,
            _ => Precedence::LOWEST,
        }
    }

    pub fn peek_precedence(&self) -> Precedence {
        if let Some(peek) = self.peek_token() {
            return match get_precedence(peek.kind) {
                Some(p) => p,
                _ => Precedence::LOWEST,
            };
        } else {
            return Precedence::LOWEST;
        }
    }

    pub fn current_token_is(&self, kind: &TokenType) -> bool {
        if let Some(peek) = self.current_token() {
            return peek.kind == *kind;
        } else {
            return false;
        }
    }

    pub fn peek_token_is(&self, kind: &TokenType) -> bool {
        if let Some(peek) = self.peek_token() {
            return peek.kind == *kind;
        } else {
            return false;
        }
    }

    pub fn expect_peek(&mut self, kind: &TokenType) -> bool {
        if variant_eq(kind, &self.peek_token().unwrap().kind) {
            self.consume_token();
            return true;
        }

        panic!(
            "{:?} Peek token was not {:?}, got {:?} instead",
            self.position,
            kind,
            self.peek_token()
        )
    }

    pub fn consume_token(&mut self) {
        self.position += 1;
    }

    pub fn new(tokens: Vec<Token>) -> Parser {
        Parser {
            tokens,
            position: 0,
        }
    }

    pub fn parseProgram(&mut self) -> Program {
        let mut statements = Vec::<Statement>::new();

        while self.position < self.tokens.len() {
            statements.push(self.parseStatement());
        }

        Program { statements }
    }

    fn parseStatement(&mut self) -> Statement {
        match self.current_token().unwrap().kind {
            // TokenType::Value => self.parseValueStatement(),
            // TokenType::Update => self.parseUpdateStatement(),
            TokenType::Return => self.parseReturnStatement(),
            _ => self.parseExpressionStatement(),
        }
    }

    fn parseReturnStatement(&mut self) -> Statement {
        let token = self.current_token().unwrap();

        self.consume_token();

        let value = self.parseExpression(Precedence::LOWEST);

        Statement::ReturnStatement { value, token }
    }

    fn parseExpressionStatement(&mut self) -> Statement {
        let token = self.current_token().unwrap();
        let expression = self.parseExpression(Precedence::LOWEST);

        if self.current_token_is(&TokenType::Semicolon) {
            self.consume_token();
        }

        Statement::ExpressionStatement { expression, token }
    }

    fn parseExpression(&mut self, precedence: Precedence) -> Box<Expression> {
        let mut left = self.matchPrefixExpression(self.current_token().unwrap().kind);

        while !self.peek_token_is(&TokenType::Semicolon)
            && int_precedence(precedence) < int_precedence(self.peek_precedence())
        {
            if let Some(peek_kind) = self.peek_token() {
                
                if self.checkInfixExpression(&peek_kind.kind) {
                    self.consume_token();
                    return self
                        .matchInfixExpression(peek_kind.kind, left.clone())
                        .unwrap();
                    
                } else {
                    break;
                }
            } else {
                println!("No infix found for {:?}", self.peek_token());
                return left;
            }
        }

        println!("Left (will be returned): {:?}", left);

        return left;
    }

    // Prefix/Infixes

    fn matchPrefixExpression(&mut self, kind: TokenType) -> Box<Expression> {
        match kind {
            TokenType::Integer(_) => self.parseIntegerExpression(),
            TokenType::Float(_) => self.parseFloatExpression(),
            TokenType::String(_) => self.parseStringExpression(),
            TokenType::Boolean(_) => self.parseBooleanExpression(),
            // TokenType::Identifier(_) => self.parseIdentifierExpression(),
            // TokenType::Minus => self.parsePrefixExpression(),
            _ => panic!(
                "[{}] PARSER ERROR: NO PREFIX FUNCTION FOUND FOR {:?}",
                self.position, kind
            ),
        }
    }

    fn parseBooleanExpression(&self) -> Box<Expression> {
        Box::new(Expression::BooleanLiteral {
            token: self.current_token().unwrap(),
        })
    }

    fn parseIntegerExpression(&self) -> Box<Expression> {
        // println!("Int");
        Box::new(Expression::IntegerLiteral {
            token: self.current_token().unwrap(),
        })
    }

    fn parseFloatExpression(&self) -> Box<Expression> {
        // println!("Int");
        Box::new(Expression::FloatLiteral {
            token: self.current_token().unwrap(),
        })
    }

    fn parseStringExpression(&self) -> Box<Expression> {
        Box::new(Expression::StringLiteral {
            token: self.current_token().unwrap(),
        })
    }

    fn checkInfixExpression(&self, kind: &TokenType) -> bool {
        match *kind {
            TokenType::Plus => true,
            TokenType::Minus => true,
            TokenType::Asterisk => true,
            TokenType::Slash => true,
            TokenType::Equal => true,
            TokenType::NotEqual => true,
            TokenType::LT => true,
            TokenType::GT => true,
            TokenType::LTEq => true,
            TokenType::GTEq => true,
            _ => false,
        }
    }

    fn matchInfixExpression(
        &mut self,
        kind: TokenType,
        leftExpression: Box<Expression>,
    ) -> Option<Box<Expression>> {
        println!("PARSING INFIX WITH LEFT: {:?}", leftExpression);
        match kind {
            TokenType::Plus => Some(self.parseInfixExpression(leftExpression)),
            TokenType::Minus => Some(self.parseInfixExpression(leftExpression)),
            TokenType::Asterisk => Some(self.parseInfixExpression(leftExpression)),
            TokenType::Slash => Some(self.parseInfixExpression(leftExpression)),
            TokenType::Equal => Some(self.parseInfixExpression(leftExpression)),
            TokenType::NotEqual => Some(self.parseInfixExpression(leftExpression)),
            TokenType::LT => Some(self.parseInfixExpression(leftExpression)),
            TokenType::GT => Some(self.parseInfixExpression(leftExpression)),
            TokenType::LTEq => Some(self.parseInfixExpression(leftExpression)),
            TokenType::GTEq => Some(self.parseInfixExpression(leftExpression)),
            _ => None,
        }
    }

    fn parseInfixExpression(&mut self, left: Box<Expression>) -> Box<Expression> {
        let tok = self.current_token().unwrap();

        let precedence = self.current_precedence();
        self.consume_token();

        println!("----PARSING RIGHT, current tok: {:?}", self.current_token());

        let right = self.parseExpression(precedence);
        self.consume_token();

        Box::new(Expression::InfixExpression {
            token: tok,

            right,
            left,
        })
    }
}
