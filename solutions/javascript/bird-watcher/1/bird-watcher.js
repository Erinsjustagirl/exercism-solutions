// @ts-check

export function totalBirdCount(birdsPerDay) {
  let count = 0;
    for(let i = 0; i < birdsPerDay.length; i++){
      count += birdsPerDay[i];
    }
  return count;
}

export function birdsInWeek(birdsPerDay, week) {
  birdsPerDay = Array.from(birdsPerDay);
  const weekDays = week * 7;
  const birdArray = birdsPerDay.slice(weekDays - 7, weekDays);
  const result = totalBirdCount(birdArray);
  return result;
}

export function fixBirdCountLog(birdsPerDay) {
    for(let i = 0; i < birdsPerDay.length; i += 2){
      birdsPerDay[i]++;
    }
}
