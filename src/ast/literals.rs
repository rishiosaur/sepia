use std::{collections::HashMap, iter::Map};

use crate::lexer::Token;

use super::{nodes::Expression, statements::BlockStatement};

pub struct IdentifierLiteral {
    token: Token,
}

impl Expression for IdentifierLiteral {
    fn literal(&self) -> String {
        todo!()
    }

    fn string(&self) -> String {
        todo!()
    }
}

pub struct StringLiteral {
    token: Token,
}

impl Expression for StringLiteral {
    fn literal(&self) -> String {
        todo!()
    }

    fn string(&self) -> String {
        todo!()
    }
}

pub struct ArrayLiteral {
    token: Token,
    elements: Vec<Box<dyn Expression>>,
}

impl Expression for ArrayLiteral {
    fn literal(&self) -> String {
        todo!()
    }

    fn string(&self) -> String {
        todo!()
    }
}

pub struct BooleanLiteral {
    token: Token,
}

impl Expression for BooleanLiteral {
    fn literal(&self) -> String {
        todo!()
    }

    fn string(&self) -> String {
        todo!()
    }
}

pub struct IndexExpression {
    token: Token,
    left: Box<dyn Expression>,
    index: Box<dyn Expression>,
}

impl Expression for IndexExpression {
    fn literal(&self) -> String {
        todo!()
    }

    fn string(&self) -> String {
        todo!()
    }
}

pub struct MapLiteral {
    token: Token,
    pairs: HashMap<Box<dyn Expression>, Box<dyn Expression>>,
}

impl Expression for MapLiteral {
    fn literal(&self) -> String {
        todo!()
    }

    fn string(&self) -> String {
        todo!()
    }
}

pub struct FunctionLiteral<'a, 'b> {
    token: Token,
    parameters: Vec<&'a IdentifierLiteral>,
    body: &'b BlockStatement,
}

impl<'a, 'b> Expression for FunctionLiteral<'a, 'b> {
    fn literal(&self) -> String {
        todo!()
    }

    fn string(&self) -> String {
        todo!()
    }
}

pub struct IntegerLiteral {
    token: Token,
}

impl Expression for IntegerLiteral {
    fn literal(&self) -> String {
        todo!()
    }

    fn string(&self) -> String {
        todo!()
    }
}

pub struct FloatLiteral {
    token: Token,
}

impl Expression for FloatLiteral {
    fn literal(&self) -> String {
        todo!()
    }

    fn string(&self) -> String {
        todo!()
    }
}
