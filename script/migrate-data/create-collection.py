from pymongo import MongoClient
# pip3 install pymongo

# Set the MongoDB connection details
MONGODB_HOST = "localhost"
MONGODB_PORT = 27017
MONGODB_USER = "root"
MONGODB_PASS = "password"
MONGODB_DB = "newsdb"
MONGODB_COLLECTION = "submissions"

try:
    # Connect to MongoDB and query the collection
    client = MongoClient('mongodb://%s:%s@%s:%d/%s' % (MONGODB_USER, MONGODB_PASS, MONGODB_HOST, MONGODB_PORT, MONGODB_DB))
    # Select or create the database
    db = client[MONGODB_DB]
    # Create the collection
    collection = db.create_collection(MONGODB_COLLECTION)
    print(f"Collection '{MONGODB_COLLECTION}' created successfully.")
except Exception as e:
    print(f"Error connecting to MongoDB: {e}")
