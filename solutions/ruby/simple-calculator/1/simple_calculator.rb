class SimpleCalculator
  ALLOWED_OPERATIONS = ['+', '/', '*'].freeze
  class UnsupportedOperation < StandardError
  end

  def self.calculate(first_operand, second_operand, operation)
    result = if !ALLOWED_OPERATIONS.include?(operation)
      raise UnsupportedOperation.new
    elsif !first_operand.is_a?(Numeric) || !second_operand.is_a?(Numeric)
      raise ArgumentError.new
    elsif operation == "+"
      first_operand + second_operand
    elsif operation == "/"
      begin
        first_operand / second_operand
      rescue ZeroDivisionError
        return "Division by zero is not allowed."
      end
    else # will be *
      first_operand * second_operand
    end
    "#{first_operand} #{operation} #{second_operand} = #{result}"
  end
end
