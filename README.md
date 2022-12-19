# regression

linear, quadratic regression

```
package regression // import "github.com/Konstantin8105/regression"


FUNCTIONS

func Linear(data [][2]float64) (
	a, b float64,
	R2 float64,
	err error,
)
    Linear regression model:

        y   = a*x+b
        R^2 - the relative predictive power of a linear model

func Quadratic(data [][2]float64) (
	a, b, c float64,
	R2 float64,
	err error,
)
    Quadratic regression model:

        y   = a*x^2+b*x+c
        R^2 - the relative predictive power of a quadratic model
```
