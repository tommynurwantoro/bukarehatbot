class CreateHistories < ActiveRecord::Migration[5.2]
  def change
    create_table :histories, id: false, primary_key: :id do |t|
      t.primary_key :id, :unsigned_integer, auto_increment: true
      t.integer :user_id, null: false
      t.references :user
      t.integer :group_id, null: false
      t.references :group
      t.integer :point, default: 0
      t.timestamp :created_at, null: false
    end
  end
end
