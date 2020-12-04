use crate::lexer::Token;

use super::{
    literals::IdentifierLiteral,
    nodes::{Expression, Statement},
};

pub struct BlockStatement {
    token: Token,
    statements: Vec<Box<dyn Statement>>,
}

impl Statement for BlockStatement {
    fn literal(&self) -> String {
        todo!()
    }

    fn string(&self) -> String {
        todo!()
    }
}

pub struct ExpressionStatement {
    token: Token,
    expression: Box<dyn Expression>,
}

impl Statement for ExpressionStatement {
    fn literal(&self) -> String {
        todo!()
    }

    fn string(&self) -> String {
        todo!()
    }
}

pub struct ValueStatement {
    token: Token,
    name: IdentifierLiteral,
    value: Box<dyn Expression>,
}

impl<'a> Statement for ValueStatement {
    fn literal(&self) -> String {
        todo!()
    }

    fn string(&self) -> String {
        todo!()
    }
}

pub struct UpdateStatement {
    token: Token,
    name: IdentifierLiteral,
    value: Box<dyn Expression>,
}

impl<'a> Statement for UpdateStatement {
    fn literal(&self) -> String {
        todo!()
    }

    fn string(&self) -> String {
        todo!()
    }
}

pub struct ReturnStatement {
    token: Token,
    value: Box<dyn Expression>,
}

impl Statement for ReturnStatement {
    fn literal(&self) -> String {
        todo!()
    }

    fn string(&self) -> String {
        todo!()
    }
}
