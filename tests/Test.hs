import Test.Hspec

main :: IO ()
main = hspec $ do
    describe "Simple statement" $ do
        it "1+1=2" $ do
            1+1 `shouldBe` 2
