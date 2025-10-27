class BoutiqueInventory
  def initialize(items)
    @items = items
  end

  def item_names
    @items.map { |n| n[:name] }.sort
  end

  def cheap
    @items.select { |n| n[:price] < 30.00}
  end

  def out_of_stock
    @items.select { |n| n[:quantity_by_size].empty? }
  end

  def stock_for_item(name)
    @items.find { |n| n[:name] == name }&.dig(:quantity_by_size) 
  end

  def total_stock
    @items.map do |n|
      n[:quantity_by_size].values.sum
    end.sum
  end

  private
  attr_reader :items
end
