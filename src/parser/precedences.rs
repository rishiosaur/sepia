use crate::lexer::TokenType;

pub enum Precedence {
    LOWEST = 0,
    AND = 1,
    OR = 2,
    EQUALS = 3,
    LESSGREATER = 4,

    SUM = 5,
    PRODUCT = 6,
    PREFIX = 7,
    CALL = 8,
    INDEX = 9
}

pub fn get_precedence(kind: TokenType) -> Option<Precedence> {
    match kind {
        TokenType::Equal => Some(Precedence::EQUALS),
        TokenType::NotEqual => Some(Precedence::EQUALS),
        TokenType::LT => Some(Precedence::LESSGREATER),
        TokenType::GT => Some(Precedence::LESSGREATER),
        TokenType::LTEq => Some(Precedence::LESSGREATER),
        TokenType::GTEq => Some(Precedence::LESSGREATER),
        TokenType::Or => Some(Precedence::OR),
        TokenType::And => Some(Precedence::AND),
        TokenType::Plus => Some(Precedence::SUM),
        TokenType::Minus => Some(Precedence::SUM),
        TokenType::Asterisk => Some(Precedence::PRODUCT),
        TokenType::Slash => Some(Precedence::PRODUCT),
        TokenType::LParen => Some(Precedence::CALL),
        TokenType::LBracket => Some(Precedence::INDEX),
        _ => None
    }
}

pub fn int_precedence (precedence: Precedence) -> usize {
    precedence as usize
}