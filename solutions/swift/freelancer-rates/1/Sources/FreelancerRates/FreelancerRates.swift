func dailyRateFrom(hourlyRate: Int) -> Double {
  let floatRate = Double(hourlyRate);
  return floatRate * 8.0;
}


// i referenced the solution from the user MaKai-Pan to complete the monthlyRateFrom function. i was just too stumped :c - Erin
func monthlyRateFrom(hourlyRate: Int, withDiscount discount: Double) -> Double {
  let monthlyRate = dailyRateFrom(hourlyRate: hourlyRate) * 22.0;
  let finalValue = monthlyRate * (1 - Double(discount) * 0.01);
  return finalValue.rounded();
}

func workdaysIn(budget: Double, hourlyRate: Int, withDiscount discount: Double) -> Double {
  let dailyRate = dailyRateFrom(hourlyRate: hourlyRate) * (1 - Double(discount) * 0.01);
  let days = (budget / dailyRate).rounded(.down);
  return days;
}
