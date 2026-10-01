// i keep pasting my code!!! yayyyy :3 - Erin
func raindrops(_ number: Int) -> String {
  var result = "";
  
  if (number % 3 == 0 ){
      result = result + "Pling";
  }
  if (number % 5 == 0) {
  	result = result + "Plang";
  }
  if (number % 7 == 0) {
  	  result = result + "Plong";
  }
  if (result == "") {
        result = String(number);
  }
    
    return result;
  }
