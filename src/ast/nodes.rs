
pub trait Node {
    fn literal(&self) -> String;
    fn string(&self) -> String;
}

pub trait Expression {
    fn literal(&self) -> String;
    fn string(&self) -> String;
    fn expressionNode(&self) {  }
}

pub trait Statement {
    fn literal(&self) -> String;
    fn string(&self) -> String;
    fn statementNode(&self) { }
}

pub struct Program {
    statements: Vec<Box<dyn Statement>>
}

impl Node for Program {
    fn literal(&self) -> String {
        if self.statements.len() != 0 {
            self.statements[0].string()
        } else {
            "".to_owned()
        }
    }

    fn string(&self) -> String {
        let mut owned = String::new();

        for node in self.statements.iter() {
            owned.push_str(node.string().as_str());
        }

        owned
    }
}

