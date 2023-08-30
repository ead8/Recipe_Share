export const GET_RECIPES = gql`
  query GetRecipes {
    recipes {
      id
      title
      description
      preparation_time
      featured_image
      category_id
      user_id
      steps
      ingredients
      images
      ratings{
        rating_value
      }
      likes {
        id
        liked
      }
      bookmarks {
        id
        bookmarked
      }
    }
  }
`;

export const GET_RECIPES_BY_TITLE = gql`
  query GetRecipesByTitle($title: String!) {
    recipes(where: { title: { _ilike: $title } }) {
      id
      title
      description
      preparation_time
      featured_image
      category_id
      user_id
      steps
      ingredients
      images
      likes {
        id
        liked
      }
      bookmarks {
        id
        bookmarked
      }
    }
  }
`;

export const GET_CATEGORIES = gql`
  query GetCategories {
    categories {
      id
      name
    }
  }
`;

export const GET_CREATORS = gql`
  query GetCreators {
    users {
      id
      username
    }
  }
`;

export const GET_RECIPES_BY_CATEGORY = gql`
  query GetRecipesByCategory($categoryId: uuid!) {
    recipes(where: { category_id: { _eq: $categoryId } }) {
      id
      title
      description
      preparation_time
      featured_image
      category_id
      user_id
      steps
      ingredients
      images
      likes {
        id
        liked
      }
      bookmarks {
        id
        bookmarked
      }
    }
  }
`;

export const GET_RECIPES_BY_CREATOR = gql`
  query GetRecipesByCreator($creatorId: uuid!) {
    recipes(where: { user_id: { _eq: $creatorId } }) {
      id
      title
      description
      preparation_time
      featured_image
      category_id
      user_id
      steps
      ingredients
      images
      likes {
        id
        liked
      }
      bookmarks {
        id
        bookmarked
      }
    }
  }
`;


export const GET_RECIPE_LIKES =gql `
query GetRecipeLikes($recipeId: uuid!) {
    likes_count: likes_aggregate(where:{
      recipe_id:{_eq:$recipeId}
    }) {
      aggregate {
        count
      }
    }
  }
`;

export const DELETE_RECIPE = gql`
  mutation DeleteRecipe($recipeId: uuid!) {
    delete_recipes_by_pk(id: $recipeId) {
      id
    }
  }
`;

export const TOGGLE_BOOKMARK =gql `
  mutation ToggleBookmark($recipeId: uuid!, $bookmarked: Boolean!) {
    insert_bookmarks(objects: { recipe_id: $recipeId, bookmarked: $bookmarked }) {
      affected_rows
    }
  }
`;
 
export const DELETE_BOOKMARK = gql`
  mutation DeleteBookmark($recipeId: uuid!) {
    delete_bookmarks(where: { recipe_id: { _eq: $recipeId } }) {
      affected_rows
    }
  }
`;

export const TOGGLE_LIKE = gql`
  mutation ToggleLike($recipeId: uuid!, $liked: Boolean!) {
    insert_likes(objects: { recipe_id: $recipeId, liked: $liked }) {
      affected_rows
    }
  }
`;

export const DELETE_LIKE = gql`
  mutation DeleteLike($recipeId: uuid!) {
    delete_likes(where: { recipe_id: { _eq: $recipeId } }) {
      affected_rows
    }
  }
`;

export const ADD_RECIPE =gql`
  mutation AddRecipe(
    $title: String!,
    $description: String,
    $preparation_time: Int,
    $featured_image: String, 
    $category_id: uuid,
    $steps: [String!]!,
    $ingredients: [String!]!,
    $images: [String!]!, 
  ) {
    insert_recipes_one(
      object: {
        title: $title,
        description: $description,
        preparation_time: $preparation_time,
        featured_image: $featured_image,
        category_id: $category_id,
        steps: $steps,
        ingredients: $ingredients,
        images: $images,
      }
    ) {
      id
    }
  }
`;

export const GET_RECIPE_BY_ID = gql`
query GetRecipeById($id: uuid!) {
  recipes_by_pk(id: $id) {
    id
    title
    description
    preparation_time
    featured_image
    category {
      id
      name
    }
    user_id  # Include the user ID here
    steps
    ingredients
    images
  }
}
`;

export const FETCH_USER_BY_ID = gql`
query GetUserById($userId: uuid!) {
  users_by_pk(id: $userId) {
    username
  }
}
`;

export const INSERT_COMMENT = gql`
mutation InsertComment($recipeId: uuid!, $commentText: String!) {
  insert_comments_one(object: {
    recipe_id: $recipeId
    comment_text: $commentText
  }) {
    id
    recipe_id
    comment_text
  }
}
`;

export const GET_SELECTED_RATING = gql`
query GetSelectedRating($recipeId: uuid!) {
  ratings(where: { recipe_id: { _eq: $recipeId } }) {
    rating_value
  }
}
`;

export const INSERT_RATING = gql`
mutation SetRecipeRating($recipeId: uuid!, $ratingValue: Int) {
  insert_ratings(objects: [{recipe_id: $recipeId, 
    rating_value: $ratingValue}]) {
    affected_rows
  }
}
`;

export const GET_BOOKMARKS = gql`
  query GetBookmarks {
    bookmarks {
      id
      user_id
      recipe_id
    }
  }
`;

export const UPDATE_RECIPE = gql`
mutation UpdateRecipe(
  $id: uuid!,
  $title: String,
  $description: String,
  $preparation_time: Int,
  $featured_image: String,
  $category_id: uuid,
  $steps: [String],
  $ingredients: [String],
  $images:[String],
) {
  update_recipes(
    where: { id: { _eq: $id } },
    _set: {
      title: $title,
      description: $description,
      preparation_time: $preparation_time,
      featured_image: $featured_image,
      category_id: $category_id,
      steps: $steps,
      ingredients: $ingredients,
      images: $images,
    }
  ) {
    affected_rows
  }
}
`;

export const GET_RECIPES_BY_PK=gql`
query GetRecipes_By_Pk($id:uuid!){
    recipes_by_pk(id:$id) {
        title
        description
        preparation_time
        featured_image
        category_id
        user_id
        steps
        ingredients
        images
    }
}
`
export const GET_RECIPES_BY_USER_ID = gql`
query GetUserRecipes($userId:uuid!) {
  recipes(where: {user_id: {_eq:$userId}}) {
    id
    title
    description
    preparation_time
    featured_image
    category_id
    user_id
    steps
    ingredients
    images
    likes {
      id
      liked
    }
    bookmarks {
      id
      bookmarked
    }
  }
}
`;