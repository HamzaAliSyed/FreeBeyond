import { MainPageData } from "../../data";
import Card from "../components/Card";

const MainPage = () => {
  return (
    <div className="container mx-auto mt-4 px-4">
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
        {MainPageData.map((card) => (
          <Card
            key={card.id}
            image={card.image}
            name={card.name}
            description={card.description}
            link={card.link}
          />
        ))}
      </div>
    </div>
  );
};
export default MainPage;
