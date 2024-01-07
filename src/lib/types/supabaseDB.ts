export type Json =
  | string
  | number
  | boolean
  | null
  | { [key: string]: Json | undefined }
  | Json[]

export interface Database {
  public: {
    Tables: {
      ingredient: {
        Row: {
          description: string | null
          id: number
          name: string
          type: Database["public"]["Enums"]["ingredient_type"]
        }
        Insert: {
          description?: string | null
          id?: number
          name: string
          type: Database["public"]["Enums"]["ingredient_type"]
        }
        Update: {
          description?: string | null
          id?: number
          name?: string
          type?: Database["public"]["Enums"]["ingredient_type"]
        }
        Relationships: []
      }
      profile: {
        Row: {
          avatar_url: string | null
          bio: string | null
          created_at: string | null
          display_name: string | null
          id: string
          onboarding_state:
            | Database["public"]["Enums"]["onboarding_state"]
            | null
        }
        Insert: {
          avatar_url?: string | null
          bio?: string | null
          created_at?: string | null
          display_name?: string | null
          id: string
          onboarding_state?:
            | Database["public"]["Enums"]["onboarding_state"]
            | null
        }
        Update: {
          avatar_url?: string | null
          bio?: string | null
          created_at?: string | null
          display_name?: string | null
          id?: string
          onboarding_state?:
            | Database["public"]["Enums"]["onboarding_state"]
            | null
        }
        Relationships: [
          {
            foreignKeyName: "profile_id_fkey"
            columns: ["id"]
            isOneToOne: true
            referencedRelation: "users"
            referencedColumns: ["id"]
          }
        ]
      }
      recipe: {
        Row: {
          brew_type: Database["public"]["Enums"]["brew_type"] | null
          created_at: string
          description: string | null
          difficulty: Database["public"]["Enums"]["difficulty"] | null
          final_gravity: number | null
          id: number
          name: string | null
          original_gravity: number | null
          process_steps: string[] | null
          published: boolean
          sweetened_gravity: number | null
          updated_at: string | null
          user_id: string | null
        }
        Insert: {
          brew_type?: Database["public"]["Enums"]["brew_type"] | null
          created_at?: string
          description?: string | null
          difficulty?: Database["public"]["Enums"]["difficulty"] | null
          final_gravity?: number | null
          id?: number
          name?: string | null
          original_gravity?: number | null
          process_steps?: string[] | null
          published?: boolean
          sweetened_gravity?: number | null
          updated_at?: string | null
          user_id?: string | null
        }
        Update: {
          brew_type?: Database["public"]["Enums"]["brew_type"] | null
          created_at?: string
          description?: string | null
          difficulty?: Database["public"]["Enums"]["difficulty"] | null
          final_gravity?: number | null
          id?: number
          name?: string | null
          original_gravity?: number | null
          process_steps?: string[] | null
          published?: boolean
          sweetened_gravity?: number | null
          updated_at?: string | null
          user_id?: string | null
        }
        Relationships: [
          {
            foreignKeyName: "recipe_user_id_fkey"
            columns: ["user_id"]
            isOneToOne: false
            referencedRelation: "profile"
            referencedColumns: ["id"]
          }
        ]
      }
      recipe_ingredient: {
        Row: {
          id: number
          ingredient_id: number | null
          quantity: number | null
          recipe_id: number
          unit: Database["public"]["Enums"]["unit_of_measurement"] | null
        }
        Insert: {
          id?: number
          ingredient_id?: number | null
          quantity?: number | null
          recipe_id: number
          unit?: Database["public"]["Enums"]["unit_of_measurement"] | null
        }
        Update: {
          id?: number
          ingredient_id?: number | null
          quantity?: number | null
          recipe_id?: number
          unit?: Database["public"]["Enums"]["unit_of_measurement"] | null
        }
        Relationships: [
          {
            foreignKeyName: "recipe_ingredient_ingredient_id_fkey"
            columns: ["ingredient_id"]
            isOneToOne: false
            referencedRelation: "ingredient"
            referencedColumns: ["id"]
          },
          {
            foreignKeyName: "recipe_ingredient_recipe_id_fkey"
            columns: ["recipe_id"]
            isOneToOne: false
            referencedRelation: "recipe"
            referencedColumns: ["id"]
          }
        ]
      }
    }
    Views: {
      [_ in never]: never
    }
    Functions: {
      [_ in never]: never
    }
    Enums: {
      brew_type:
        | "Ale"
        | "Lager"
        | "Stout"
        | "IPA"
        | "Mead"
        | "Melomel"
        | "Cyser"
        | "Hydromel"
        | "Metheglin"
        | "Cider"
        | "Fruit Wine"
        | "Other"
      difficulty: "Easy" | "Intermediate" | "Hard"
      ingredient_type:
        | "Grain"
        | "Hops"
        | "Yeast"
        | "Fruit"
        | "Spice"
        | "Honey"
        | "Sugar"
        | "Nutrient"
        | "Additives"
        | "Other"
      onboarding_state:
        | "email_unconfirmed"
        | "display_name_pending"
        | "bio_pending"
        | "avatar_pending"
        | "completed"
      unit_of_measurement:
        | "g"
        | "kg"
        | "oz"
        | "lb"
        | "ml"
        | "l"
        | "tsp"
        | "tbsp"
        | "cup"
        | "pint"
        | "quart"
        | "gal"
    }
    CompositeTypes: {
      [_ in never]: never
    }
  }
}

export type Tables<
  PublicTableNameOrOptions extends
    | keyof (Database["public"]["Tables"] & Database["public"]["Views"])
    | { schema: keyof Database },
  TableName extends PublicTableNameOrOptions extends { schema: keyof Database }
    ? keyof (Database[PublicTableNameOrOptions["schema"]]["Tables"] &
        Database[PublicTableNameOrOptions["schema"]]["Views"])
    : never = never
> = PublicTableNameOrOptions extends { schema: keyof Database }
  ? (Database[PublicTableNameOrOptions["schema"]]["Tables"] &
      Database[PublicTableNameOrOptions["schema"]]["Views"])[TableName] extends {
      Row: infer R
    }
    ? R
    : never
  : PublicTableNameOrOptions extends keyof (Database["public"]["Tables"] &
      Database["public"]["Views"])
  ? (Database["public"]["Tables"] &
      Database["public"]["Views"])[PublicTableNameOrOptions] extends {
      Row: infer R
    }
    ? R
    : never
  : never

export type TablesInsert<
  PublicTableNameOrOptions extends
    | keyof Database["public"]["Tables"]
    | { schema: keyof Database },
  TableName extends PublicTableNameOrOptions extends { schema: keyof Database }
    ? keyof Database[PublicTableNameOrOptions["schema"]]["Tables"]
    : never = never
> = PublicTableNameOrOptions extends { schema: keyof Database }
  ? Database[PublicTableNameOrOptions["schema"]]["Tables"][TableName] extends {
      Insert: infer I
    }
    ? I
    : never
  : PublicTableNameOrOptions extends keyof Database["public"]["Tables"]
  ? Database["public"]["Tables"][PublicTableNameOrOptions] extends {
      Insert: infer I
    }
    ? I
    : never
  : never

export type TablesUpdate<
  PublicTableNameOrOptions extends
    | keyof Database["public"]["Tables"]
    | { schema: keyof Database },
  TableName extends PublicTableNameOrOptions extends { schema: keyof Database }
    ? keyof Database[PublicTableNameOrOptions["schema"]]["Tables"]
    : never = never
> = PublicTableNameOrOptions extends { schema: keyof Database }
  ? Database[PublicTableNameOrOptions["schema"]]["Tables"][TableName] extends {
      Update: infer U
    }
    ? U
    : never
  : PublicTableNameOrOptions extends keyof Database["public"]["Tables"]
  ? Database["public"]["Tables"][PublicTableNameOrOptions] extends {
      Update: infer U
    }
    ? U
    : never
  : never

export type Enums<
  PublicEnumNameOrOptions extends
    | keyof Database["public"]["Enums"]
    | { schema: keyof Database },
  EnumName extends PublicEnumNameOrOptions extends { schema: keyof Database }
    ? keyof Database[PublicEnumNameOrOptions["schema"]]["Enums"]
    : never = never
> = PublicEnumNameOrOptions extends { schema: keyof Database }
  ? Database[PublicEnumNameOrOptions["schema"]]["Enums"][EnumName]
  : PublicEnumNameOrOptions extends keyof Database["public"]["Enums"]
  ? Database["public"]["Enums"][PublicEnumNameOrOptions]
  : never
