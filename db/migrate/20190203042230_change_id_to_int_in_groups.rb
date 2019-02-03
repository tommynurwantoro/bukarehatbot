class ChangeIdToIntInGroups < ActiveRecord::Migration[5.2]
  def change
    change_column :groups, :id, :int
  end
end
