// use parser::Parser;

mod ast;
mod evaluator;
mod lexer;
mod objects;
mod parser; 
mod util;

fn main() {
    let lex = lexer::Lexer::new(
        r#"
        1 * 2 + 4"#,
    );

    let z: Vec<lexer::Token> = lex.collect();
    println!("{:#?}", z);

    // let mut p = Parser::new(z);
    // println!("{:#?}", p.parseProgram());
}
