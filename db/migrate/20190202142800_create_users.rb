class CreateUsers < ActiveRecord::Migration[5.2]
  def change
    create_table :users, id: false, primary_key: :id do |t|
      t.primary_key :id, :unsigned_integer, auto_increment: true
      t.string :username, null: false
      t.references :group
      t.boolean :is_admin, null: false
      t.integer :point, default: 0

      t.timestamps
    end
  end
end
