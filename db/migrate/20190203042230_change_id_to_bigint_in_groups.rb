class ChangeIdToBigintInGroups < ActiveRecord::Migration[5.2]
  def change
    change_column :groups, :id, :bigint
  end
end
