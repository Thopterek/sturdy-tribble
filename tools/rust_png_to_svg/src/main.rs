use image::{GenericImageView, Pixel, Rgba};

fn main() {
    let path_to_img = "../../overview/pixel_art_sketches/bird.png";
    let res = image::open(path_to_img);
    if let Ok(img) = res {
        let dimensions = img.dimensions();
        println!("Dimensions {:?}", img.dimensions());
        if dimensions.0 != 16 && dimensions.1 != 16 {
            println!("Error: Wrong image my guy, here we work on 16 by 16");
            println!("Info: Should it be an error? No. Do I care? Not enough");
            return;
        }
        for one in img.pixels() {
            let r = one.2;
            r.to_rgb();
            // r.map();
        }
    } else {
        println!("Error: It uses relative path to find image {}", path_to_img);
    }
}
