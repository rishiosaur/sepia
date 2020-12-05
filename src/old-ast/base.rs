use crate::lexer::{Token, TokenType};

use crate::ast::nodes::Expression;

pub struct PrefixExpression {
    token: Token,
    right: Box<dyn Expression>,
}

impl Expression for PrefixExpression {
    fn literal(&self) -> String {
        format!("{:?}", self.token)
    }
    fn string(&self) -> String {
        format!("({:?}{:?})", self.token.position, self.token)
    }
}

pub struct InfixExpression {
    token: Token,
    right: Box<dyn Expression>,
    left: Box<dyn Expression>,
}

impl Expression for InfixExpression {
    fn literal(&self) -> String {
        format!("{:?}", self.token)
    }
    fn string(&self) -> String {
        format!("({:?}{:?})", self.token.position, self.token)
    }
}
