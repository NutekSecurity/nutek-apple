# frozen_string_literal: true

# @return [Boolean] true if the user's response was "yes" or "y", false otherwise
# @param prompt [String] The prompt to display to the user
def get_yes_no_input(prompt)
  # Print the prompt to the console
  print prompt

  # Read the user's input from standard input (STDIN)
  answer = $stdin.gets

  # Remove any newline characters from the input and convert it to lowercase
  answer = answer.chomp.downcase

  # Return true if the input matches "yes" or "y", false otherwise
  return true if %w[yes y].include?(answer)

  # Otherwise, return false
  false
end
