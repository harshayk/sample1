
package main




import (


	"html/template"

	

"net/http"


	"strconv"
 
      
)







func main() {

	
var n1,n2,n3 int64

	
var a,b string

	
var i int

	
tmpl := template.Must(template.New("calc1").Parse(`
<html>

<head>

<title> Calculator</title>

</head>

<body >

<center>

<form name="calculator" method="POST" >

<input type="textfield" name="ans" style="height:50px; font-size: 18px" >

<br />

<textarea name="body" rows="3" cols="18" style="font-size: 18px" >{{ . }}
</textarea>
<br />

<input type="button" value="1" style="height:50px; width:50px; font-size: 18px"onClick="document.calculator.ans.value+='1'">

<input type="button" value="2" style="height:50px; width:50px; font-size: 18px"onClick="document.calculator.ans.value+='2'">

<input type="button" value="3" style="height:50px; width:50px; font-size: 18px"onClick="document.calculator.ans.value+='3'"><br />

<input type="button" value="4" style="height:50px; width:50px; font-size: 18px"onClick="document.calculator.ans.value+='4'">

<input type="button" value="5" style="height:50px; width:50px; font-size: 18px"onClick="document.calculator.ans.value+='5'">

<input type="button" value="6" style="height:50px; width:50px; font-size: 18px"onClick="document.calculator.ans.value+='6'"><br />

<input type="button" value="7" style="height:50px; width:50px; font-size: 18px"onClick="document.calculator.ans.value+='7'">

<input type="button" value="8" style="height:50px; width:50px; font-size: 18px"onClick="document.calculator.ans.value+='8'">

<input type="button" value="9" style="height:50px; width:50px; font-size: 18px"onClick="document.calculator.ans.value+='9'"><br />

<input type="button" value="0" style="height:50px; width:158px; font-size: 18px"onClick="document.calculator.ans.value+='0'"><br />

<input type="button" value="-" style="height:50px; width:50px; font-size: 18px"onClick="document.calculator.ans.value+='-'">

<input type="button" value="+" style="height:50px; width:50px; font-size: 18px"onClick="document.calculator.ans.value+='+'">

<input type="button" value="*" style="height:50px; width:50px; font-size: 18px"onClick="document.calculator.ans.value+='*'"><br />

<input type="button" value="%" style="height:50px; width:50px; font-size: 18px"onClick="document.calculator.ans.value+='%'">

<input type="button" value="/" style="height:50px; width:50px; font-size: 18px"onClick="document.calculator.ans.value+='/'">
<input type="reset" value="C" style="width:50px; height:50px">

<br /><input type="submit" value="=" style="height:50px; width:158px">      

</form>

</center>

</body>

</html>

`))


	
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {


		if r.Method != http.MethodPost {


			tmpl.Execute(w, nil)


			return

	
	        }

	              s := r.FormValue("ans")


                       
           for i=0; i<len(s); i++  {



                    if s[i]=='+' || s[i]=='-' || s[i]=='*' || s[i]=='/' || s[i]=='%'  {

                          a = s[0:i]

                       
   b = s[i+1:len(s)]


		          n1, _ = strconv.ParseInt(a, 10, 64)


		          n2, _ = strconv.ParseInt(b, 10, 64)


		       	
  break


	 
              }
 
          
         } 


		
	
				if i!=len(s)  {

	                        	if s[i]=='+' {

                                             n3 = n1+n2

                                             a = strconv.FormatInt(n3, 10) 
           
      	             }  else if s[i]=='-' {

                                             n3 = n1-n2

                        	             a = strconv.FormatInt(n3, 10)

	                             }  else if s[i]=='*' {
                  
                           n3 = n1*n2

                                             a = strconv.FormatInt(n3, 10)

                                     }  else if s[i]=='/' {

                                             if n2!=0  {
                         
			    n3 = n1/n2

		                                    a = strconv.FormatInt(n3, 10)

                                              }  else   {a = "division by zero not possible"}

                                     }  else if s[i]=='%' {

		                             n3 = n1%n2

                       			     a = strconv.FormatInt(n3, 10)

	                             }  
				
             
                               } else  {
                      
 	             n3, _ = strconv.ParseInt(s, 10, 64)

                                     a = strconv.FormatInt(n3, 10)	
		
                                 }

         
      
	
	                      
                                tmpl.Execute(w, a)

	
                    

    
	  })

	

                               http.ListenAndServe(":8080", nil)
 

}
