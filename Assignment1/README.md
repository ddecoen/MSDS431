# Assignment 1: Testing Go for Statistics

### The Anscombe Quartet
### Go testing package
### Go statistical package
### https://github.com/montanaflynn/stats

## Can Go be used for statistics in a similar fashion to Python and R
The answer is yes it can, but it may still be better to perform higher level analysis in Python or R. 

While Go is very fast, the actual coding took much longer as compared to Python or R.   
For example, Python has an extensive library that has several built in commands to use with basic statistics such as Linear Regression.  
When attempting to investigate Anscombe's Quartet in Golang, it was necessary to create several functions to use in conjunction with the stats package.  

Python's statsmodels.api had a built in model for Ordinary Least Squares ("OLS"). It was easy to just run the model in Python as well as print the results.   
Leveraging matplotlib.pyplot also allowed for friendly data visualizations. R is even better with visualizations. This is an area for Golang to continue to improve upon.  

![Anscombe Quartet] (~/images/fig_anscombe_Python.png)






#### demonstration data from
#### Anscombe, F. J. 1973, February. Graphs in statistical analysis. 
####  The American Statistician 27: 17–21.
