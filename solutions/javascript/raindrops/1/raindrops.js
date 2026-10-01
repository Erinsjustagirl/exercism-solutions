// i copied some code from my go one because its roughly the same - Erin

export const convert = (number) => {
  let result = "";

  if (number % 3 == 0) {
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
