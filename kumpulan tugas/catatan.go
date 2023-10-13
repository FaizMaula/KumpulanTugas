Chandelure 5428 249 134 735 110 24.23
Cinderace 4724 339 193 91 116 103.39
Cramorant 5589 250 243 576 162 40.53 
Decidueye 4546 359 193 114 116 89.02 
Delphox 5448 248 134 705 108 20.77 
Dragapult 4796 377 187 87 126 60.29 


/* Binary Search

untuk data ascending
func cari(s pokemon,n int, s string){
	var kr,kn,mid int
	var found bool = false
	
	kr = 0
	kn = n-1
	for kr <= kn && !found {
		mid = (kr+kn)/2
		if T[mid] > s {
			kn = mid - 1
		} else if T[mid] < s {
			kr = mid + 1
		} else {
			found = true
		}
		
	}

}

Sequential Search

func cari(T pokemon, n int, s string) int {
	var j int
	for i := 0 ; i<n ;i++{
		if T[i] == s {
			j = i
		}
	}
	return j
	OR
	var found int = -1
	for i < n && found = -1 {
		if T[i] == s {
			found = i 
		}
	}
	return found
	
}

Insertion Sort

PSEUDOCODE

DESCENDING
DESCENDING
DESCENDING
procedure selsort(in/out t : array, in n : integer)
kamus
	j,i integer
	temp array
algoritma
	while i < n-1 do
		j <- i
		temp <- T[j]
		while j > 0 && temp > T[j-1] do // tuker temp < T[j-1]
			T[j] <- T[j-1]
			j <- j - 1
		endwhile
		T[j] <- temp
		i <- i + i
	endwhile
endprocudere
GoLang	
func insSort(T* tPokemon, n int) {
	var j int
	var temp pokemon
	for i := 0;i<n-1;i++ {
		j = i
		temp = T[j]
		for j > 0 && temp.SP_Def > T[j-1].SP_Def {
			T[j] = T[j-1]
			j--
		}
		T[j] = temp
	}
}

ASCENDING
ASCENDING
ASCENDING
SELECTION SORT
PSEUDOCODE
procedure selSort(in/out tPokemon, in n integer) 
kamus 
	idx,i,pass : integer
	temp : pokemon
algoritma
	pass <- 1
	while pass < n do
		idx <- pass - 1
		i <- 0 
		while i < n do
			if T[idx] > T[i] then // T[idx] < T[i] DESCENDING
				idx <- i
			endif
			i <- i + 1
		endwhile
		temp <- T[pass-1}
		T[pass-1} <- T[idx]
		T[idx] <- temp
	endwhile
endprocedure
		
	
GoLang

func selsort(T* tPokemon,n int) {
	var pass,idx,i int
	var temp pokemon
	for pass := 1;pass<n-1;pass++{
		idx = pass - 1
		for i:=0;i<n;i++{
			if T[idx] > T[i] {
				idx = i
			}
		}
		temp = T[pass-1]
		T[pass-1] = T[idx]
		T[idx] = temp
	}
}
*/



procedure selSort(in/out s tPokemon, n integer) 

kamus 
	pass,i,idx : integer
	temp : pokemon
algoritma
	pass <- 1
	while pass < n do
		idx <- pass -1
		for i <- 0 to n do 
			if T[idx] <T[i] then
				idx <- i
			endif
		endfor
		temp <- T[pass-1]
		T[pass-1] <- T[idx]
		T[idx] <- temp
		pass <- pass + 1
		
endprocedure

procedure insSort(in/out s tPokemon, n integer)
kamus
	j :integer
	temp : pokemon
algoritma
	while i < n do
		j = i
		temp = T[j]
		for j > 0 && temp > T[j-1] do
			T[j] = T[j-1]
			j <- j - 1
		endfor
		T[j] = temp
	endwhile