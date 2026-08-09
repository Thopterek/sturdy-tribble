pub fn get_color(color: u8) -> String {
    if color < 40 {
        String::from("\\outline\\")
    } else if color > 40 && color < 80 {
        String::from("\\darkspot\\")
    } else if color > 80 && color < 118 {
        String::from("\\shadow\\")
    } else if color > 125 && color < 180 {
        String::from("\\slight\\")
    } else if color > 180 && color < 220 {
        String::from("\\gbase\\")
    } else {
        String::from("\\white\\")
    }
}

pub fn get_symbol_for_terminal(color: u8) -> String {
    if color < 40 {
        String::from("#")
    } else if color > 40 && color < 80 {
        String::from("X")
    } else if color > 80 && color < 118 {
        String::from("$")
    } else if color > 125 && color < 180 {
        String::from("Z")
    } else if color > 180 && color < 220 {
        String::from("U")
    } else {
        String::from(" ")
    }
}
