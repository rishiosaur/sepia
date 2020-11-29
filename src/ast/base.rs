use crate::lexer::{Token, TokenType};

use crate::ast::nodes::{Expression};

pub struct PrefixExpression {
    token: Token,
    operator: TokenType,
    right: Box<dyn Expression>
}

impl Expression for PrefixExpression {
    fn literal(&self) -> String { format!("{:?}", self.token) }
    fn string(&self) -> String { format!("({:?}{:?})", self.operator, self.token) }
}

pub struct InfixExpression {
    token: Token,
    operator: TokenType,
    right: Box<dyn Expression>
}

impl Expression for InfixExpression {
    fn literal(&self) -> String { format!("{:?}", self.token) }
    fn string(&self) -> String { format!("({:?}{:?})", self.operator, self.token) }
}

