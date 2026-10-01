let birthday = "Birthday";
let valentine = "Valentine\'s Day";
let anniversary = "Anniversary";

let space: Character = " ";
let exclamation: Character = "!";

func buildSign(for occasion: String, name: String) -> String {
  let resultingSign = "Happy \(occasion) \(name)!";
  return resultingSign;
}

func graduationFor(name: String, year: Int) -> String {
  let gradResult = "Congratulations \(name)!\nClass of \(year)";
  return gradResult;
}

func costOf(sign: String) -> Int {
  let baseCost = 20;
  let cost = (sign.count * 2) + baseCost;
  return cost;
}
