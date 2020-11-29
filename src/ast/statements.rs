use crate::lexer::Token;

use super::{literals::Identifier, nodes::{Expression, Statement}};

pub struct BlockStatement {
    token: Token,
    statements: Vec<Box<dyn Statement>>
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
    expression: Box<dyn Expression>
}

impl Statement for ExpressionStatement {
    fn literal(&self) -> String {
        todo!()
    }

    fn string(&self) -> String {
        todo!()
    }
}

pub struct ValueStatement<'a> {
    token: Token,
    name: &'a Identifier,
    value: Box<dyn Expression>
}

impl<'a> Statement for ValueStatement<'a> {
    fn literal(&self) -> String {
        todo!()
    }

    fn string(&self) -> String {
        todo!()
    }
}

pub struct UpdateStatement<'a> {
    token: Token,
    name: &'a Identifier,
    value: Box<dyn Expression>
}

impl<'a> Statement for UpdateStatement<'a> {
    fn literal(&self) -> String {
        todo!()
    }

    fn string(&self) -> String {
        todo!()
    }
}


pub struct ReturnStatement {
    token: Token,
    value: Box<dyn Expression>
}

impl Statement for ReturnStatement {
    fn literal(&self) -> String {
        todo!()
    }

    fn string(&self) -> String {
        todo!()
    }
}
