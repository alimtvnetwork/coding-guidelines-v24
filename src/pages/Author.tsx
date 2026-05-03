import { Link } from "react-router-dom";
import { ArrowLeft, Linkedin } from "lucide-react";
import { Button } from "@/components/ui/button";

export default function Author() {
  return (
    <main className="min-h-screen bg-background text-foreground">
      <section className="container mx-auto max-w-3xl px-6 py-16">
        <Button asChild variant="ghost" size="sm" className="mb-8">
          <Link to="/">
            <ArrowLeft className="mr-2 h-4 w-4" aria-hidden="true" />
            Back home
          </Link>
        </Button>

        <h1 className="font-heading text-4xl font-bold tracking-tight sm:text-5xl">
          Meet <span className="text-gradient-brand">Md. Alim Ul Karim</span>
        </h1>
        <p className="mt-3 text-base text-muted-foreground">
          Software Architect · Chief Software Engineer at{" "}
          <a
            href="https://riseup-asia.com/"
            target="_blank"
            rel="noopener noreferrer"
            className="text-primary underline-offset-4 hover:underline"
          >
            Riseup Asia LLC
          </a>
        </p>

        <div className="prose-spec mt-8 space-y-5 text-base leading-relaxed text-foreground/90">
          <p>
            Md. Alim Ul Karim started his programming journey in <strong>2004</strong> — not
            for school or a job, but because video games made him deeply curious about how
            computers really worked. While most beginners write "Hello, World", his very
            first program was an attempt to build his own database engine. By <strong>2005</strong>,
            frustrated with writing raw SQL, he wrote his own{" "}
            <a
              href="https://clink.rasia.pro/alim-orm-2005-vb"
              target="_blank"
              rel="noopener noreferrer"
              className="text-primary underline-offset-4 hover:underline"
            >
              ORM
            </a>
            , becoming one of the earliest .NET adopters in his country.
          </p>
          <p>
            Alim has authored multiple production frameworks and earned recognition as a{" "}
            <strong>Top 1% talent</strong> by joining and working with crossover.com.
          </p>
        </div>

        <div className="mt-8">
          <Button asChild variant="outline" className="lift-hover">
            <a
              href="https://www.linkedin.com/in/alimkarim/"
              target="_blank"
              rel="noopener noreferrer"
            >
              <Linkedin className="mr-2 h-4 w-4" aria-hidden="true" />
              Connect on LinkedIn
            </a>
          </Button>
        </div>

        <section className="mt-12 space-y-4">
          <h2 className="font-heading text-2xl font-semibold">Alim believes in</h2>
          <blockquote className="border-l-4 border-primary/60 bg-card/60 px-5 py-4 text-lg italic text-foreground/90">
            "You don't need to start with the complex stuff. You need to start with curiosity."
          </blockquote>
          <blockquote className="border-l-4 border-primary/60 bg-card/60 px-5 py-4 text-lg italic text-foreground/90">
            "Pain creates principles."
          </blockquote>
          <p className="text-right text-sm text-muted-foreground">— Md. Alim Ul Karim</p>
        </section>
      </section>
    </main>
  );
}
