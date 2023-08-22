Maps are the Key-value pair in Golang in a format of map[key]value.

Key can be only of specific data types but Value can be of any type.
for Reference to what data types are allowed as key in maps than visit [w3schools.com](https://www.w3schools.com/go/go_maps.php)

We can even iterate over maps like we did on Slices but with For-Range only to get key-value as individual in Go.

We can delcare map variable in different ways:
    -var mapVariableName map[keyDataType]ValueDataType 
    -using make() function eg. make(map[keyDataType]ValueDataType)