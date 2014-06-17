# image2json

jordic@gmail.com
http://tempointeractiu.net

image2json, takes a small image, an converts it, to json string, representing, positions and their colors. 

## Install 

- Clone the repo
- go build image2json.go

## Flags
``` image2json
 -f="": Image file to decode
  -nc=false: Output image without color info
  -o="": file to output
```

f: Input image
o: output file
nc: Output mode without colors

The tool can ouput two diferent json strings, first with info color:
(without -nc flag)

```js
{
    "Width":93,
    "Height":11,
    "Bytes":[{"X":1,"Y":2,"C":{"R":91,"G":184,"B":255,"A":255}}, ... }
```

Where Bytes holds and array of color positions, X,Y and their color balue in RGB. **White color on image rgb(0,0,0) is ommited**.

Second mode ( without color info ) flag -nc:

```js
{   "Width":13,
    "Height":12,
    "Bytes":[[2,0],[3,0],[4,0]... }
```

Where byts only, outputs, position with color. For better results, this option must be used with single image colors. Also, **white** color is ommited


## For What I can use such tool

I builded it for deploying javascript "particle" animations, see examples folder.

```
  ***   ***  
 *   * *   * 
*     *     *
*           *
*           *
*           *
 *         * 
  *       *  
   *     *   
    *   *    
     * *     
      *      
```

