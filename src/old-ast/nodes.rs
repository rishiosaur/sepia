use std::collections::HashMap;

use crate::lexer::Token;

#[derive(Debug, Clone)]
pub struct Program<'a> {
    statements: Vec<Box<&'a Statement<'a>>>,
}
#[derive(Debug, Clone)]
pub enum Expression<'a> {
    IdentifierLiteral {
        token: &'a Token,
    },
    StringLiteral {
        token: &'a Token,
    },
    ArrayLiteral {
        token: &'a Token,
        elements: &'a Vec<Box<Expression<'a>>>,
    },
    BooleanLiteral {
        token: &'a Token,
    },

    FloatLiteral {
        token: &'a Token,
    },
    IndexExpression {
        token: &'a Token,
        left: &'a Box<Expression<'a>>,
        index: &'a Box<Expression<'a>>,
    },
    MapLiteral {
        token: &'a Token,
        pairs: HashMap<&'a Box<Expression<'a>>, &'a Box<Expression<'a>>>,
    },
    FunctionLiteral {
        token: &'a Token,
        parameters: &'a Vec<Expression<'a>>,
    },

    PrefixExpression {
        token: &'a Token,
        right: &'a Box<Expression<'a>>,
    },

    InfixExpression {
        token: &'a Token,
        right: &'a Box<Expression<'a>>,
        left: &'a Box<Expression<'a>>,
    },

    IfExpression {
        token: Token,
        condition: &'a Box<dyn Expression>,
        consequence: &'a Statement<'a>,
        alternative: &'a Statement<'a>,
    },
}

#[derive(Debug, Clone)]
pub enum Statement<'a> {
    BlockStatement {
        token: &'a Token,
        statements: Vec<&'a Box<Statement<'a>>>,
    },
    ExpressionStatement {
        token: &'a Token,
        statements: Vec<&'a Box<Statement<'a>>>,
    },
    ValueStatement {
        token: &'a Token,
        name: &'a Expression<'a>,
        value: &'a Box<Expression<'a>>,
    },
    UpdateStatement {
        token: &'a Token,
        name: &'a Expression<'a>,
        value: &'a Box<Expression<'a>>,
    },
    ReturnStatement {
        token: &'a Token,
        value: &'a Box<Expression<'a>>,
    },
}
