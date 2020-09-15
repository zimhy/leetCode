package main

import (
	"sort"
)

func findItinerary(tickets [][]string) []string {

	allPlacesOrderMap := make(map[string]int)
	allPlaces := make([]string, 0)
	for i := 0; i < len(tickets); i++ {
		allPlacesOrderMap[tickets[i][0]] = 0
		allPlacesOrderMap[tickets[i][1]] = 0
	}
	for key := range allPlacesOrderMap {
		allPlaces = append(allPlaces, key)
	}
	sort.Strings(allPlaces)
	for i := 0; i < len(allPlaces); i++ {
		allPlacesOrderMap[allPlaces[i]] = i
	}
	ticketStatusMap := make([][]int, len(allPlaces))
	for i := 0; i < len(allPlaces); i++ {
		ticketStatusMap[i] = make([]int, len(allPlaces))
		for j := 0; j < len(allPlaces); j++ {
			ticketStatusMap[i][j] = 0
		}
	}
	for i := 0; i < len(tickets); i++ {
		ticketStatusMap[allPlacesOrderMap[tickets[i][0]]][allPlacesOrderMap[tickets[i][1]]]++
	}
	itinerary := make([]string, 0)
	itinerary = append(itinerary, "JFK")
	from, ok := allPlacesOrderMap["JFK"]
	if !ok {
		return itinerary
	}
	ok = findItineraryFromMap(from, &ticketStatusMap, &itinerary, &allPlaces, len(tickets))
	if ok {
		return itinerary
	} else {
		return nil
	}
}

func findItineraryFromMap(from int, ticketStatusMap *[][]int, itinerary *[]string, allPlaces *[]string, ticketsCount int) bool {
	if len(*itinerary) == (ticketsCount + 1) {
		return true
	}
	for to, status := range (*ticketStatusMap)[from] {
		if status > 0 {
			(*ticketStatusMap)[from][to] = (*ticketStatusMap)[from][to] - 1
			itineraryNew := append(*itinerary, (*allPlaces)[to])
			ok := findItineraryFromMap(to, ticketStatusMap, &itineraryNew, allPlaces, ticketsCount)
			if ok {
				*itinerary = itineraryNew
				return true
			} else {
				(*ticketStatusMap)[from][to] = (*ticketStatusMap)[from][to] + 1
			}
		}
	}
	return false
}

//func main() {
//	var decoded =
//	data := "[[\"AUA\",\"PER\"],[\"LST\",\"ADL\"],[\"CNS\",\"TIA\"],[\"ADL\",\"VIE\"],[\"ADL\",\"VIE\"],[\"BNE\",\"CBR\"],[\"EZE\",\"VIE\"],[\"JFK\",\"ADL\"],[\"CBR\",\"HBA\"],[\"CNS\",\"AUA\"],[\"HBA\",\"BNE\"],[\"OOL\",\"LST\"],[\"PER\",\"AUA\"],[\"SYD\",\"AXA\"],[\"TIA\",\"BNE\"],[\"MEL\",\"AXA\"],[\"AUA\",\"OOL\"],[\"LST\",\"OOL\"],[\"DRW\",\"SYD\"],[\"CNS\",\"SYD\"],[\"INN\",\"CBR\"],[\"BNE\",\"INN\"],[\"BNE\",\"EZE\"],[\"BNE\",\"CNS\"],[\"OOL\",\"DRW\"],[\"BNE\",\"EZE\"],[\"CBR\",\"BNE\"],[\"TIA\",\"LST\"],[\"OOL\",\"JFK\"],[\"SYD\",\"CBR\"],[\"PER\",\"MEL\"],[\"HBA\",\"OOL\"],[\"MEL\",\"EZE\"],[\"OOL\",\"HBA\"],[\"AUA\",\"PER\"],[\"DRW\",\"HBA\"],[\"VIE\",\"ANU\"],[\"HBA\",\"BNE\"],[\"DRW\",\"TIA\"],[\"AXA\",\"VIE\"],[\"LST\",\"BNE\"],[\"CNS\",\"MEL\"],[\"ADL\",\"HBA\"],[\"VIE\",\"OOL\"],[\"TIA\",\"MEL\"],[\"PER\",\"DRW\"],[\"INN\",\"CNS\"],[\"JFK\",\"LST\"],[\"LST\",\"DRW\"],[\"MEL\",\"TIA\"],[\"EZE\",\"CNS\"],[\"AXA\",\"CNS\"],[\"ADL\",\"LST\"],[\"TIA\",\"JFK\"],[\"VIE\",\"SYD\"],[\"INN\",\"JFK\"],[\"VIE\",\"ADL\"],[\"SYD\",\"AUA\"],[\"ANU\",\"INN\"],[\"BNE\",\"SYD\"],[\"JFK\",\"INN\"],[\"SYD\",\"PER\"],[\"ADL\",\"TIA\"],[\"JFK\",\"ADL\"],[\"CBR\",\"ADL\"],[\"EZE\",\"BNE\"]]"
//	json.Unmarshal([]byte(data), &decoded)
//	[["MUC","LHR"],["JFK","MUC"],["SFO","SJC"],["LHR","SFO"]]
//	var tickets = [][]string{
//		[]string{"MUC","LHR"},
//		[]string{"JFK","MUC"},
//		[]string{"SFO","SJC"},
//		[]string{"LHR","SFO"},
//		//[]string{"MUC", "ABC"},
//		//[]string{"LHR", "SFO"},
//	}
//	var result = findItinerary(decoded)
//	fmt.Println("%v", result)
//}
