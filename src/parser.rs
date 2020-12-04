use crate::{lexer::Token, util::variant_eq};

use crate::lexer::TokenType;

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
    pub fn current_token(&self) -> Option<&Token> {
        self.tokens.get(self.position)
    }

    pub fn peek_token(&self) -> Option<&Token> {
        self.tokens.get(self.position + 1)
    }

    pub fn current_precedence(&self) -> Precedence {
        match get_precedence(self.current_token().unwrap().clone().kind) {
            Some(p) => p,
            _ => Precedence::LOWEST,
        }
    }

    pub fn peek_precedence(&self) -> Precedence {
        match get_precedence(self.peek_token().unwrap().clone().kind) {
            Some(p) => p,
            _ => Precedence::LOWEST,
        }
    }

    pub fn current_token_is(&self, kind: &TokenType) -> bool {
        self.current_token().unwrap().kind == *kind
    }

    pub fn peek_token_is(&self, kind: &TokenType) -> bool {
        self.peek_token().unwrap().kind == *kind
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

    pub fn new() -> Parser {

    }
}
