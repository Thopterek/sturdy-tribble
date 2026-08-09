use image::{GenericImageView, Pixel};

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
            let access_rgba = one.2;
            let actuall_rgb_pixel = access_rgba.to_rgb();
            let r = actuall_rgb_pixel.0;
            println!("{:?}", r);
        }
    } else {
        println!("Error: It uses relative path to find image {}", path_to_img);
    }
}
