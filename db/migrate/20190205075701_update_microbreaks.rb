class UpdateMicrobreaks < ActiveRecord::Migration[5.2]
  def change
    remove_column :microbreaks, :name
    remove_column :microbreaks, :rest_time
    add_column :microbreaks, :rest_minute, :integer, limit: 2, after: :url
    add_column :microbreaks, :rest_hour, :integer, limit: 2, after: :url
  end
end
