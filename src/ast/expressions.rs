use crate::lexer::Token;

use super::{statements::BlockStatement, nodes::Expression};

struct IfExpression<'a> {
    token: Token,
    condition: Box<dyn Expression>,
    consequence: &'a BlockStatement,
    alternative: &'a BlockStatement
}

impl<'a> Expression for IfExpression<'a> {
    fn literal(&self) -> String {
        todo!()
    }

    fn string(&self) -> String {
        todo!()
    }
}

pub struct CallExpression {
    token: Token,
    function: Box<dyn Expression>,
    arguments: Vec<Box<dyn Expression>>
}

impl Expression for CallExpression {
    fn literal(&self) -> String {
        todo!()
    }

    fn string(&self) -> String {
        todo!()
    }
}

