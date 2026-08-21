use image::GenericImageView;

use rust_png_to_svg::print_png_on_terminal;

/*
* Just checking if I am pushing with correct user and from new branch
*/
fn main() {
    let path_to_img = "../../overview/pixel_art_sketches/bird.png";
    let res = image::open(path_to_img);
    if let Ok(img) = res {
        let dimensions = img.dimensions();
        println!("Dimensions {:?}", dimensions);
        if dimensions.0 != 16 && dimensions.1 != 16 {
            println!("Error: Wrong image my guy, here we work on 16 by 16");
            println!("Info: Should it be an error? No. Do I care? Not enough");
            return;
        }
        print_png_on_terminal(img);
    } else {
        println!("Error: It uses relative path to find image {}", path_to_img);
    }
}
