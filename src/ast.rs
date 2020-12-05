use std::collections::HashMap;

use crate::lexer::Token;

#[derive(Debug, Clone)]
pub struct Program {
    pub statements: Vec<Statement>,
}
#[derive(Debug, Clone)]
pub enum Expression {
    IdentifierLiteral {
        token: Token,
    },
    StringLiteral {
        token: Token,
    },
    ArrayLiteral {
        token: Token,
        elements: Vec<Expression>,
    },
    BooleanLiteral {
        token: Token,
    },
    IntegerLiteral {
        token: Token,
    },
    FloatLiteral {
        
        token: Token,
    },
    IndexExpression {
        token: Token,
        left: Box<Expression>,
        index: Box<Expression>,
    },
    MapLiteral {
        token:  Token,
        pairs: HashMap<Box<Expression>,Box<Expression>>,
    },
    FunctionLiteral {
        token:  Token,
        parameters:  Vec<Expression>,
    },

    PrefixExpression {
        token:  Token,
        right:  Box<Expression>,
    },

    InfixExpression {
        token:  Token,
        right:  Box<Expression>,
        left:  Box<Expression>,
    },

    IfExpression {
        token: Token,
        condition:  Box<Expression>,
        consequence:  Statement,
        alternative:  Statement,
    },
}

#[derive(Debug, Clone)]
pub enum Statement {
    BlockStatement {
        token:  Token,
        statements: Vec<Statement>,
    },
    ExpressionStatement {
        token:  Token,
        expression: Box<Expression>,
    },
    ValueStatement {
        token:  Token,
        name:  Token,
        value:  Box<Expression>,
    },
    UpdateStatement {
        token:  Token,
        name:  Token,
        value:  Box<Expression>,
    },
    ReturnStatement {
        token:  Token,
        value:  Box<Expression>,
    },
}
