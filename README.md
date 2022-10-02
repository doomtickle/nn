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
