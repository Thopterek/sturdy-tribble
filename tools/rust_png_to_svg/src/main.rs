use std::fs::read;

fn main() {
    println!("Hello, we are changing this dummy from Haskell!");
    let mut file = read("../../../overview/pixel_art_sketches/bird.png");
    // println!(file);
}
