// i pasted my code again. its just so fun to port the same 15~20 lines of code to different languages - Erin
pub fn raindrops(n: u32) -> String {
    let mut result = String::new();
  
    if n % 3 == 0 {
      result = result + "Pling";
    }
    if n % 5 == 0 {
  	result = result + "Plang";
    }
    if n % 7 == 0 {
  	  result = result + "Plong";
    }
    if result == "" {
        result = n.to_string();
    }
    
    return result;
}
