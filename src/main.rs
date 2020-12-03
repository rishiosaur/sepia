use ast::base::PrefixExpression;

mod ast;
mod evaluator;
mod lexer;
mod objects;
mod parser;
mod util;

fn main() {
    let lex = lexer::Lexer::new(
        r#"
value z = "hello waef";
value x = 3.2444;
print(z + x);
"#,
    );

    let z: Vec<lexer::Token> = lex.collect();
    println!("{:#?}", z);
}
