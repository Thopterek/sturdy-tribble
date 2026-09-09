use image::{DynamicImage, GenericImageView, Pixel};

#[cfg(debug_assertions)]
pub fn get_color(color: u8) -> String {
    if color < 1 {
        String::from("SKIP")
    } else if color < 40 {
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

#[cfg(debug_assertions)]
pub fn print_png_on_terminal(img: DynamicImage) {
    for one in img.pixels().enumerate() {
        let access_rgba = one.1.2;
        let actuall_rgb_pixel = access_rgba.to_rgba();
        // let result = get_color(actuall_rgb_pixel.0[0]);
        if actuall_rgb_pixel.0[3] <= 128 {
            print!(" ");
        } else {
            let result = get_symbol_for_terminal(actuall_rgb_pixel.0[0]);
            print!("{}", result);
        }
        if (one.0 + 1) % 16 == 0 {
            println!();
        }
    }
}

pub fn c_to_char(color: u8) -> char {
    if color <= 1 {
        'B'
    } else if (80..=100).contains(&color) {
        'D'
    } else if (118..=140).contains(&color) {
        'S'
    } else if (150..=170).contains(&color) {
        'L'
    } else if (200..220).contains(&color) {
        'B'
    } else {
        'W'
    }
}

#[cfg(debug_assertions)]
pub fn get_symbol_for_terminal(color: u8) -> String {
    if color <= 1 {
        String::from("B")
    } else if (80..=100).contains(&color) {
        String::from("%")
    } else if (118..=140).contains(&color) {
        String::from("#")
    } else if (150..=170).contains(&color) {
        String::from("\"")
    } else if (200..220).contains(&color) {
        String::from("\'")
    } else {
        String::from("W")
    }
}

#[cfg(debug_assertions)]
pub fn print_rgba(img: DynamicImage) {
    for one in img.pixels().enumerate() {
        let access_rgba = one.1.2;
        let actuall_rgb_pixel = access_rgba.to_rgba();
        // let result = get_color(actuall_rgb_pixel.0[0]);
        println!("{:?}", actuall_rgb_pixel);
        if (one.0 + 1) % 16 == 0 {
            println!();
        }
    }
}
