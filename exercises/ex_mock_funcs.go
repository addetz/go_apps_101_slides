//     func TestSomethingThatUsesDB(t *testing.T) {
//         // make and configure a mocked DB
//         mockedDB := &DBMock{
//             FindOrderFunc: func(id string) (models.Order, error) {
// 	               panic("mock out the FindOrder method")
//             },
//         }