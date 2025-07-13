func validMountainArray(arr []int) bool {
    i := 1

    for i < len(arr) && arr[i] > arr[i-1] { // checks arr bounds and ensures ascending
        i++
    }

    if i == 1 || i == len(arr) { return false } // arr length not valid for a mountain array

    for i < len(arr) && arr[i] < arr[i-1] { // checks arr bounds and ensures descending
        i++
    }

    return i == len(arr)
}