mod ast;
mod evaluator;
mod lexer;
mod objects;
mod parser;
mod util;

fn main() {
    let lex = lexer::Lexer::new(
        r#"value z = "hello
waef" 3.45"#);

    let z: Vec<lexer::Token> = lex.collect();
    println!("{:#?}", z);
    
}
