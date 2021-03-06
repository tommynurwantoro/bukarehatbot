# This file is auto-generated from the current state of the database. Instead
# of editing this file, please use the migrations feature of Active Record to
# incrementally modify your database, and then regenerate this schema definition.
#
# Note that this schema.rb definition is the authoritative source for your
# database schema. If you need to create the application database on another
# system, you should be using db:schema:load, not running all the migrations
# from scratch. The latter is a flawed and unsustainable approach (the more migrations
# you'll amass, the slower it'll run and the greater likelihood for issues).
#
# It's strongly recommended that you check this file into your version control system.

ActiveRecord::Schema.define(version: 2019_02_07_043213) do

  create_table "break_histories", id: :integer, unsigned: true, options: "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci", force: :cascade do |t|
    t.bigint "user_id"
    t.bigint "group_id"
    t.integer "point", default: 0
    t.timestamp "created_at", default: -> { "CURRENT_TIMESTAMP" }, null: false
    t.index ["group_id"], name: "index_break_histories_on_group_id"
    t.index ["user_id"], name: "index_break_histories_on_user_id"
  end

  create_table "groups", id: :bigint, default: nil, options: "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci", force: :cascade do |t|
    t.string "name", null: false
  end

  create_table "microbreaks", id: :integer, unsigned: true, options: "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci", force: :cascade do |t|
    t.bigint "group_id"
    t.string "url", null: false
    t.integer "rest_hour", limit: 2
    t.integer "rest_minute", limit: 2
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
    t.index ["group_id"], name: "index_microbreaks_on_group_id"
  end

  create_table "users", id: :integer, unsigned: true, options: "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci", force: :cascade do |t|
    t.string "username", null: false
    t.bigint "group_id"
    t.boolean "is_admin", null: false
    t.integer "point", default: 0
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
    t.index ["group_id"], name: "index_users_on_group_id"
  end

end
