import unittest
from app import app, db
from models import Category, Subcategory, Product

class TestApp(unittest.TestCase):
    
    def setUp(self):
        self.app = app.test_client()
        self.app.testing = True
        self.ctx = app.app_context()
        self.ctx.push()
        db.create_all()

    def tearDown(self):
        db.session.remove()
        db.drop_all()
        self.ctx.pop()

    def test_get_all_categories(self):
        # Add sample categories to the database for testing
        category1 = Category(name="Category 1")
        category2 = Category(name="Category 2")
        db.session.add(category1)
        db.session.add(category2)
        db.session.commit()

        response = self.app.get('/categories')
        data = response.get_json()

        self.assertEqual(response.status_code, 200)
        self.assertEqual(len(data['categories']), 2)

    def test_get_all_subcategories(self):
        # Add sample subcategories to the database for testing
        subcategory1 = Subcategory(name="Subcategory 1")
        subcategory2 = Subcategory(name="Subcategory 2")
        db.session.add(subcategory1)
        db.session.add(subcategory2)
        db.session.commit()

        response = self.app.get('/subcategories')
        data = response.get_json()

        self.assertEqual(response.status_code, 200)
        self.assertEqual(len(data['subcategories']), 2)

if __name__ == '__main__':
    unittest.main()
