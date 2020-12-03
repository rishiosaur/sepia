use crate::{
    ast::nodes::Program,
    lexer::{Lexer, Token, TokenType},
    util::variant_eq,
};

use super::precedences::{get_precedence, Precedence};

pub struct Parser<'a> {
    pub tokens: &'a Vec<Token>,
    pub position: usize,
}

impl<'a> Parser<'a> {
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
        match get_precedence(self.peek_token().unwrap().clone().kind) {
            Some(p) => p,
            _ => Precedence::LOWEST,
        }
    }

    pub fn current_token_is(&self, kind: TokenType) -> bool {
        self.current_token().unwrap().kind == kind
    }

    pub fn peek_token_is(&self, kind: TokenType) -> bool {
        self.peek_token().unwrap().kind == kind
    }

    pub fn expect_peek(&self, kind: TokenType) -> bool {
        if variant_eq(kind, self.peek_token().unwrap().kind) {
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
}
