# Fibo Series : 1,1,2,3,5,8,13,21,34,55,...
def fibonacci(n):
    if n <= 1:
        return n
    else:
        return fibonacci(n - 1) + fibonacci(n - 2)

# Example usage:
result = fibonacci(10)
print(result)
