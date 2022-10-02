Toy Neural Network Library in Go.

## Multi Layered Perceptron

`inputs -> [?] -> outputs`

Out Range bounds = -1 -> 1

Supervised learning

What happens inside the neuron?

Missing pieces...(Weights and biases)

Perceptron creates a sum of all inputs multiplied by their weights.

Example: 2 in -> perceptron -> 2 out

1. `sum(X(0) * W(0) + X(1) * W(1))`
2. Activation function (map to output range) (Sigmoid, sign, tanh, etc)

This is sometimes referred to as a "Feed Forward Neural Network"

Perceptron Algorithm

1. For every input multiply that input by its weight.
2. Sum all of the weighted inputs.
3. Compute the output of the perceptron based on that sum passed through an activation function

Supervised Learning Algorithm

1. Provide the perceptron with inputs for which there is a known answer.
2. Ask the perceptron to guess an answer.
3. Comput the error. (Did it get the answer right or wrong?)
4. Adjust all the weights according to the error.
5. GOTO 1

∆W = `Error / Input`
New Weight = `Weight + Error * Input * Learning Rate`

Calculate using "Gradient Descent"
