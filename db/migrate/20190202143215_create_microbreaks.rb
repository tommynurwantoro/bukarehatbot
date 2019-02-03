class CreateMicrobreaks < ActiveRecord::Migration[5.2]
  def change
    create_table :microbreaks, id: false, primary_key: :id do |t|
      t.primary_key :id, :unsigned_integer, auto_increment: true
      t.references :group
      t.string :name, null: false
      t.timestamp :rest_time, null: false
      t.string :url, null: false


      t.timestamps
    end
  end
end
