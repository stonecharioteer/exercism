class AssemblyLine
  def initialize(speed)
    @speed = speed
  end

  def production_rate_per_hour
    if (1..4).include?(@speed)
      @speed * 221
    elsif (5..8).include?(@speed)
      @speed * 221 * 0.9
    elsif @speed == 9
      @speed * 221 * 0.8
    else
      @speed * 221 * 0.77
    end
  end

  def working_items_per_minute
    (production_rate_per_hour / 60).to_i
  end
end
