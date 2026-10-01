// copied code from my javascript this was sooooo easy once i figured it out - Erin
class Raindrops {
  String convert(int number) {
    String result = "";
  
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
        result = number.toString();
    }
    
    return result;
  }
}
