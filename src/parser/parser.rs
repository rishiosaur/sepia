use crate::{ast::nodes::Program, lexer::{Lexer, Token, TokenType}};

use super::precedences::{Precedence, get_precedence};

struct Parser<'a> {
    tokens: &'a Vec<Token>,
    position: usize,
}

impl<'a> Parser<'a> {
    fn current_token(&self) -> Option<Token> {
        self.tokens.get(self.position).cloned()
    }

    fn peek_token(&self) -> Option<Token> {
        self.tokens.get(self.position +1).cloned()
    }

    fn current_precedence(&self) -> Precedence {
        match get_precedence(self.current_token().unwrap().clone().kind) {
            Some(p) => p,
            _ => Precedence::LOWEST
        }
    }

    fn peek_precedence(&self) -> Precedence {
        match get_precedence(self.peek_token().unwrap().clone().kind) {
            Some(p) => p,
            _ => Precedence::LOWEST
        }
    }

    fn current_token_is(&self, kind: &TokenType) -> bool {
        self.current_token().unwrap().kind == *kind
    }

    fn peek_token_is(&self, kind: &TokenType) -> bool {
        self.peek_token().unwrap().kind == *kind
    }

    fn expect_peek(&self, kind: TokenType) -> bool {
        if self.peek_token_is(&kind) {
            return true;
        }

        panic!("{:?} Peek token was not {:?}, got {:?} instead", self.position, kind, self.peek_token())
    }

    fn consume_token(&mut self) {
        self.position += 1;
    }
}